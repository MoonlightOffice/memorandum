import { createSignal, onCleanup, onMount } from "solid-js";
import { css } from "@emotion/css";

export function CheckEffectMount() {
  const [show, setShow] = createSignal(false);
  return (
    <div>
      <button type="button" on:click={() => setShow((b) => !b)}>toggle</button>
      {show() && <Block />}
    </div>
  );
}

export function Block() {
  onMount(() => {
    console.log("Mounted");
  });

  onCleanup(() => {
    console.log("Cleaned");
  });

  return (
    <p
      class={css({
        border: "1px solid red",
      })}
    >
      Hehehe
    </p>
  );
}
