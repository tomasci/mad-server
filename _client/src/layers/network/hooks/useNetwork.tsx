import { useCallback } from "react";
import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { xfetch } from "@/src/layers/network/xfetch.ts";
import {
  EndpointParam,
  XEndpoint,
  XRequest,
  XResponse,
} from "@/src/layers/network/types.ts";
import { closeSnackbar, enqueueSnackbar } from "notistack";

const useNetwork = () => {
  // props
  const debug = useDebug();

  // state

  // functions
  const query = useCallback(async () => {
    return new Promise((resolve, reject) => {
      let attempt = 0;
      const maxAttempts = 5;

      const tryQuery = () => {
        setTimeout(() => {
          console.log(`Attempt ${attempt + 1}: useNetwork query`);
          attempt++;
          if (attempt < maxAttempts) {
            tryQuery();
          } else {
            resolve({ someResult: "here" });
          }
        }, 1000);
      };

      tryQuery();
    });
  }, []);

  const mutation = useCallback(
    async <
      Params extends EndpointParam[],
      Body = never,
      ResponseDataType = never,
    >(
      endpoint: XEndpoint<Params>,
      request: XRequest<Body>,
    ) => {
      // return new Promise((resolve, reject) => {
      try {
        // make request using fetch
        const result = await xfetch(endpoint, request);

        // parse response
        const response = (await result.json()) as XResponse<ResponseDataType>;

        // if everything fine - return response
        if (result.ok) {
          return response;
        } else {
          // else, start with checking response error
          if (response.error) {
            if (response.message) {
              // throw message from response
              throw new Error(response.message);
            }

            // if no message in response - unknown error
            throw new Error("network_unknown_error");
          }

          // if status is not error, yet it is still failed - then fatal
          throw new Error("network_fatal_error");
        }
      } catch (e: unknown) {
        debug.error("error", e);

        // // handle errors here (example of global error notification)
        // const err = e as { message: string };
        // const id = Math.random();
        // enqueueSnackbar(
        //   <>
        //     {err.message}{" "}
        //     <button
        //       type={"button"}
        //       onClick={() => {
        //         closeSnackbar(id);
        //       }}
        //     >
        //       close
        //     </button>
        //   </>,
        //   {
        //     key: id,
        //     variant: "error",
        //     hideIconVariant: true,
        //     persist: true,
        //   },
        // );

        throw e; // send error to next level
      }
    },
    [debug],
  );

  return {
    query,
    mutation,
  };
};

export { useNetwork };
