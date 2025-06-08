import type { RouteSectionProps } from "@solidjs/router";

export default function Page(props: RouteSectionProps) {
  const lang = props.params["lang"];
  return (
    <div>
      <p>Language: {lang}</p>
      <a href={`/${lang}/about`}>Go to about</a>
    </div>
  );
}
