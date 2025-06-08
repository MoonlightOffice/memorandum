import type { RouteSectionProps } from "@solidjs/router";

export default function Page(props: RouteSectionProps) {
  return <p>About: {props.params["lang"]}</p>;
}
