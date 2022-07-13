import { RouteLocationRaw } from "vue-router";

export namespace routing {
  export interface DetailOpts {
    id: number;
  }
  export function detail(params: DetailOpts): RouteLocationRaw {
    return {
      name: "detail",
      params: {
        id: params.id,
      },
    };
  }
}
