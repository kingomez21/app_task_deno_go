import type { Handlers, PageProps } from "$fresh/server.ts";
import ListView from "../components/ListView.tsx";
import Forms from "../islands/Forms.tsx";

export const handler: Handlers = {
  async GET(_req, ctx) {
    const res = await fetch("http://localhost:4000/users")
    const data = await res.json()
        
    const resp = await ctx.render(data);
    return resp;
  },
};

export default function Home(props: PageProps) {
  
  return (
    <div className="m-10 justify-center">
      <Forms />
      <br />
      <br />
      <div height={500} style={{overflow: "auto"}} >
      <ListView people={props.data} />
      </div>
      
    </div>
  );
}
