import { For } from "solid-js";
import { createAsync } from "@solidjs/router";

//export const route = {
//  preload: () => getUsers(),
//};

function getUsers(): Promise<string[]> {
  return new Promise((res) => {
    setTimeout(() => {
      res(["user 1", "user 2"]);
    }, 500);
  });
}

export default function Page() {
  const users = createAsync(() => getUsers());

  return <For each={users()}>{(user) => <li>{user}</li>}</For>;
}
