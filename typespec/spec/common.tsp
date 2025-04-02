import "@typespec/http";

using TypeSpec.Http;

model Err {
  code: string;
  msg: string;
}

model ErrUnknown extends Err {
  code: "UNKNOWN";
}

model ErrCommon extends Err {
  code: "UNKNOWN" | "NOT_FOUND" | "INVALID";
}

model Result<T, E extends Err = ErrUnknown> {
  err: E | null;
  val: T | null;
}

model Req<T> {
  @body body: T;
}

model ReqSecure<T> extends Req<T> {
  /**
   * Authorization: Bearer \<YOUR_API_TOKEN\>
   */
  @header Authorization: string;
}

model Resp<T, E extends Err = ErrUnknown> {
  @statusCode statusCode: 200;
  @body result: Result<T, E>;
}
