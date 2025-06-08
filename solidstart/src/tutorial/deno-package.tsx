import { stringify } from "@std/yaml";

export function ShowYaml() {
  const data = {
    foo: "FOO",
    bar: "BAR",
    nested: {
      level: 1,
    },
  };

  return (
    <textarea placeholder={stringify(data)} value={stringify(data)} rows={5} />
  );
}
