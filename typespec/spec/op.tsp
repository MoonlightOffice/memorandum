import "@typespec/http";
import "@typespec/sse";
import "./common.tsp";
import "@typespec/events";

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
  @get
  @summary("Get a pet record")
  op getPet(@path id: string): Resp<Pet, ErrCommon>;

  /**
   * ### Sample Documentation
   *
   * Hi, are you reading this documentation? Great!
   *
   * hehe
   * hehe
   */
  @route("/:id/create-pet")
  @post
  op createPet(req: ReqSecure<Pet>): {
    @body resp: bytes;
  };
}

@route("/experiment")
namespace Experiment {
  @TypeSpec.Events.events
  union Events {
    someEvent: SomeEvent,
    anotherEvent: AnotherEvent,

    @TypeSpec.Events.contentType("text/plain")
    @TypeSpec.SSE.terminalEvent
    "[unsubscribe]",
  }

  model SomeEvent {
    @statusCode statusCode: 200;
    err: Err | null;
    eventId: string;
    timestamp: int64;
    content: string;
  }

  model AnotherEvent {
    @statusCode statusCode: 200;
    err: Err | null;
    eventId: string;
    timestamp: int64;
    userId: string;
    content: string;
  }

  /**
   * Sample Server-Sent Event endpoint
   */
  @route("/sse")
  @post
  op sse(req: ReqSecure<null>): TypeSpec.SSE.SSEStream<Events>;

  alias Events2 = SomeEvent | AnotherEvent;

  /**
   * Sample Server-Sent Event endpoint 2
   */
  @route("/sse2")
  @post
  op sse2(req: ReqSecure<null>): Events2;
}
