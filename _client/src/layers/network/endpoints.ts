import {
  Endpoint,
  EndpointsController,
} from "@/src/layers/network/EndpointsController.ts";

const endpoints = new EndpointsController()
  .add(
    new Endpoint("users-create")
      .name("Users - Create user")
      .api("v1")
      .service("users")
      .path("create"),
  )
  .add(
    new Endpoint("users-login")
      .name("Users - Login")
      .api("v1")
      .service("users")
      .path("login"),
  );

export { endpoints };
