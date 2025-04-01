// Documentation: https://gorm.io/

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db := InitDB()

	err := db.AutoMigrate(
		Company{},
		Employee{},
	)
	if err != nil {
		panic(err)
	}

	db.Save([]Company{C1, C2})

	var comps []Company
	db.Debug().Where("? < created_at", 10000).Where("created_at < 2743481367").Find(&comps)
	println(len(comps))

	//db.Migrator().DropColumn(&Company{}, "EstablishedAt")

	//FulltextSearch(db)
}

type Company struct {
	CompanyId     string `gorm:"primaryKey;size:128"`
	Name          string `gorm:"not null;index:idx_sample,priority:1"`
	EstablishedAt int64  `gorm:"not null;index:idx_sample,priority:2"`
	CreatedAt     int64  `gorm:"not null;autoUpdateTime:false;index:created_at"`
	UpdatedAt     int64  `gorm:"not null;autoUpdateTime:false"`
}

func NewCompany(name string) Company {
	now := generateNow()
	return Company{
		CompanyId:     generateId("comp"),
		Name:          name,
		EstablishedAt: now,
		UpdatedAt:     now,
	}
}

// TableName Implementation Tabler to change the default table name
func (Company) TableName() string {
	return "COMPANY"
}

type Employee struct {
	CompanyId  string         `gorm:"primaryKey;size:128"`
	EmployeeId string         `gorm:"primaryKey;size:128"`
	Name       string         `gorm:"not null;type:longtext;index:,class:FULLTEXT,option:WITH PARSER ngram"`
	JoinedAt   int64          `gorm:"not null;index"`
	Metadata   datatypes.JSON `gorm:"not null"`
}

func NewEmployee(companyId, name string) Employee {
	b, _ := json.Marshal(map[string]any{
		"foo": map[string]any{
			"bar": "bee",
		},
	})

	return Employee{
		CompanyId:  companyId,
		EmployeeId: generateId("emp"),
		Name:       name,
		JoinedAt:   generateNow(),
		Metadata:   b,
	}
}

var (
	C1 = NewCompany("MoonlightOffice")
	C2 = NewCompany("Soleil")
	E1 = NewEmployee(C1.CompanyId, "神崎美月")
	E2 = NewEmployee(C1.CompanyId, "Mikuru Natsuki")
	E3 = NewEmployee(C2.CompanyId, "Ichigo Hoshimiya")
)

func Find(db *gorm.DB) {
	result := db.Create([]Employee{E1, E2})
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create([]Employee{E1, E2})
	if !errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		panic("expected duplicated")
	}

	// Get multiple records at once
	var ee []Employee
	result = db.Where("company_id = ?", C1.CompanyId).Limit(1).Order("name DESC").Find(&ee)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(len(ee))

	// Get multiple records one by one
	rows, err := db.Model(Employee{}).Where("company_id = ?", C1.CompanyId).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		err = db.ScanRows(rows, &e)
		if err != nil {
			panic(err)
		}
		fmt.Println("rows.scan:", e.Name)
	}

	// SELECT FOR UPDATE
	var e []Employee
	result = db.
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		Find(&e, datatypes.JSONQuery("metadata").Equals("bee", "foo", "bar"))
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(e)
}

func Update(db *gorm.DB) {
	// UPSERT
	result := db.Save(E1)
	if result.Error != nil {
		panic(result.Error)
	}

	// UPDATE
	E3 := Employee(E1)
	E3.Name = "Ichigo Hoshimiya"
	E3.JoinedAt = 0

	result = db.
		Where("employee_id = ?", E3.EmployeeId).
		Save(E3) // db.Update will only update non-zero values, so db.Save is preferable in most cases
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("RowsAffected:", result.RowsAffected)

	var e Employee
	result = db.Where("employee_id = ?", E3.EmployeeId).Find(&e)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(E1)
	fmt.Println(E3)
	fmt.Println(e)
}

func Delete(db *gorm.DB) {
	result := db.Save([]Employee{E1, E2})
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Delete(&E1)
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Where("employee_id = ?", E2.EmployeeId).Delete(Employee{})
	if result.Error != nil {
		panic(result.Error)
	}

	var e Employee
	result = db.Where("employee_id = ?", E2.EmployeeId).First(&e)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func Join(db *gorm.DB) {
	if err := db.Save([]Company{C1, C2}).Error; err != nil {
		panic(err)
	}
	if err := db.Save([]Employee{E1, E2, E3}).Error; err != nil {
		panic(err)
	}

	var employees []Employee
	db.Joins("JOIN COMPANY ON COMPANY.company_id = employees.company_id").
		Where("employees.company_id = ?", E1.CompanyId).
		Where("employees.name LIKE CONCAT(?, '%')", E1.Name).
		Find(&employees)
	fmt.Println(len(employees))
}

func FulltextSearch(db *gorm.DB) {
	db.Save([]Employee{E1, E2})

	var e Employee
	result := db.Where("MATCH (name) AGAINST ('+美月' IN BOOLEAN MODE)").First(&e)
	fmt.Println(e, result.Error)
}

func Transaction(db *gorm.DB) {
	db.Save(C1)
	fmt.Println(C1.CompanyId)

	db.Transaction(func(tx *gorm.DB) error {
		var c Company
		result := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
			Where("company_id = ?", C1.CompanyId).
			Find(&c)
		if result.Error != nil {
			return result.Error
		}

		c.Name += "-WM"
		tx.Save(c)

		return nil
	})

	var c Company
	db.Where("company_id = ?", C1.CompanyId).Find(&c)
	fmt.Println(c.Name)
}

func CheckPool(db *gorm.DB) {
	// SHOW GLOBAL STATUS LIKE 'connections';
	// SHOW PROCESSLIST
	wg := sync.WaitGroup{}
	for i := range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
				var x int
				raws := tx.Raw("SELECT 1115")
				if raws.Error != nil {
					return raws.Error
				}
				fmt.Println(i, "point 1")

				time.Sleep(time.Second)

				raws = tx.Raw("SELECT 1115")
				if raws.Error != nil {
					return raws.Error
				}
				fmt.Println(i, "point 2")

				if scanResult := raws.Scan(&x); scanResult.Error != nil {
					return scanResult.Error
				}
				fmt.Printf("check %d: %d\n", i, x)

				return nil
			})
			if err != nil {
				fmt.Println(i, err)
			}
		}()
	}
	wg.Wait()
}
