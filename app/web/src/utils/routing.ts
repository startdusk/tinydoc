import { RouteLocationRaw } from "vue-router";

export namespace routing {
  export interface DetailOpts {
    id: number;
  }
  export function detail(params: DetailOpts): RouteLocationRaw {
    return {
      name: "details",
      params: {
        id: params.id,
      },
    };
  }

  export function home(): RouteLocationRaw {
    return {
      name: "home"
    }
  }
}
