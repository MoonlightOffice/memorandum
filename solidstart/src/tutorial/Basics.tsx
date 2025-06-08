import {
  createContext,
  createEffect,
  createResource,
  createSignal,
  For,
  JSX,
  Match,
  mergeProps,
  onCleanup,
  Show,
  Suspense,
  Switch,
  useContext,
} from "solid-js";
import { createAsync } from "@solidjs/router";

export function ViewProps() {
  const [age, setAge] = createSignal(22);

  return (
    <div>
      <button type="button" on:click={() => setAge((a) => a + 1)}>
        Increment
      </button>

      <Props name="Ichigo" age={age()} />
    </div>
  );
}

export function Props(props: {
  name: string;
  age?: number;
}) {
  const mprops = mergeProps({ age: 20 }, props);

  return (
    <div>
      <Show
        when={mprops.age >= 18}
        fallback={
          <p>Hello, I'm child, {mprops.name}, {mprops.age} years old!</p>
        }
      >
        <p>Hello, I'm adult, {mprops.name}, {mprops.age} years old!</p>
      </Show>
    </div>
  );
}

export function Cond() {
  const [count, setCount] = createSignal(0);

  function handleIncrement() {
    setCount((c) => c + 1);
  }

  return (
    <div>
      <button type="button" on:click={handleIncrement}>Increment</button>
      <Switch fallback={<p>{count()} does not have reminder.</p>}>
        <Match when={count() % 3 == 1}>
          <p>{count()} has reminder 1.</p>
        </Match>
        <Match when={count() % 3 == 2}>
          <p>{count()} has reminder 2.</p>
        </Match>
      </Switch>
    </div>
  );
}

export function RenderChild(props: {
  children: JSX.Element;
  color?: string;
}) {
  const mprops = mergeProps({ color: "red" }, props);

  return (
    <div style={{ border: `1px solid ${mprops.color}` }}>
      {props.children}
    </div>
  );
}

function fetchData(): Promise<string> {
  return new Promise((res) => {
    setTimeout(() => {
      res(`Hello world at ${new Date().getTime()}`);
    }, 1000);
  });
}

export function Async() {
  const data = createAsync(fetchData);

  return (
    <div>
      <p>Hehe</p>
      <Suspense fallback={<p>Now loading...</p>}>
        <p>Result: {data()}</p>
      </Suspense>
    </div>
  );
}

export function Async1() {
  const [data, { refetch }] = createResource(fetchData);

  const timer = setInterval(() => {
    refetch();
  }, 3000);
  onCleanup(() => clearInterval(timer));

  function handleOutput() {
    setTimeout(() => {
      console.log("Hello");
    }, 1000);
  }

  return (
    <div>
      <button type="button" on:click={handleOutput}>Hello</button>
      <Suspense fallback={<p>Now loading...</p>}>
        <p>Result: {data()}</p>
      </Suspense>
    </div>
  );
}

export function Async2() {
  const [data, { refetch }] = createResource(fetchData);

  const timer = setInterval(refetch, 3000);
  onCleanup(() => clearInterval(timer));

  const [text, setText] = createSignal(
    (
      <p>Now loading...</p>
    ),
  );

  createEffect(() => {
    setText(data());
  });

  return (
    <div>
      <Suspense>
        <p>Result: {text()}</p>
      </Suspense>
    </div>
  );
}

export function Async3() {
  const [data, setData] = createSignal<string | null>(null);

  fetchData().then((d) => {
    setData(d);
  });

  return (
    <div>
      <p>Hello</p>
      <Show when={data() === null}>
        <p>Now loading...</p>
      </Show>
      <Show when={data() !== null}>
        <p>{data()}</p>
      </Show>
    </div>
  );
}

export function ListRendering() {
  const names = ["Ichigo", "Aoi", "Ran"];

  return (
    <For each={names}>
      {(item, index) => <p>Name: {item}, index: {index()}</p>}
    </For>
  );
}

export function RefSample() {
  let paragraph!: HTMLParagraphElement;

  createEffect(() => {
    console.log(paragraph.innerText);
  });

  return <p ref={paragraph}>Hi, paragraph!</p>;
}

const SomeContext = createContext();

export function ContextParent(props: { children: JSX.Element }) {
  return (
    <SomeContext.Provider value="Some value">
      <p>This is Parent</p>
      {props.children}
    </SomeContext.Provider>
  );
}

export function ContextChild() {
  const value = useContext(SomeContext);

  return <p>Child accepts value: {value as string}</p>;
}

export function ContextView() {
  return (
    <ContextParent>
      <ContextChild />
    </ContextParent>
  );
}
