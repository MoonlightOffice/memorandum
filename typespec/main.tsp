import "@typespec/http";
import "./spec/op.tsp";

using TypeSpec.Http;

@service(#{ title: "Pet Store", version: "0.1.0" })
@server("https://example.com", "Production server endpoint")
@server("https://dev.example.com", "Dev server endpoint")
namespace PetStore;
