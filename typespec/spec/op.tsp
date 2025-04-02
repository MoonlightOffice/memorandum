import "@typespec/http";
import "./common.tsp";

using TypeSpec.Http;

namespace PetStore;

model Pet {
  id: int32 = 1115;

  // Internal comment
  @minLength(1)
  name: string;

  /**
   * ## Age⚡︎
   * The age of the pet is given here
   */
  @minValue(0)
  @maxValue(150)
  age: int32;

  kind: PetType;
}

enum PetType {
  dog: "dog",
  cat: "cat",
  fish: "fish",
  bird: "bird",
  reptile: "reptile",
}

@route("/pets")
namespace Pets {
  #deprecated ""
  @summary("Get a pet record")
  @get op getPet(@path id: string): Resp<Pet, ErrCommon>;

  /**
   * ### Sample Documentation
   *
   * Hi, are you reading this documentation? Great!
   *
   * hehe
   * hehe
   */
  @post
  @route("/:id/create-pet")
  op createPet(req: ReqSecure<Pet>): Resp<Pet>;
}
