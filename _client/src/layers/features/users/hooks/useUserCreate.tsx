import { useCallback } from "react";
import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { User, User_Create_Input } from "@/src/layers/features/users/types.ts";
import { useNetwork } from "@/src/layers/network/hooks/useNetwork.tsx";
import { useAtom } from "jotai/index";
import {
  localUserNetworkErrorAtom,
  localUserStatusAtom,
  localUserUpdaterAtom,
} from "@/src/layers/features/users/store.ts";

const useUserCreate = () => {
  // props
  const debug = useDebug();
  const { mutation } = useNetwork();

  // state
  const [, setLocalUser] = useAtom(localUserUpdaterAtom);
  const [, setLocalUserStatus] = useAtom(localUserStatusAtom);
  const [, setError] = useAtom(localUserNetworkErrorAtom);

  // functions
  const userCreate = useCallback(
    async (input: User_Create_Input) => {
      debug.log("userCreate call")

      // first update all possible local states
      const updateLocalData = async () => {
        // reset error (from previous tries)
        setError(null)
        // and set local user data (you already have some from input)
        setLocalUser({
          username: input.username,
          email: input.email,
        });
      };

      // and simultaneously send request
      const performRequest = async () => {
        // set loading status
        setLocalUserStatus("loading");

        try {
          // mutate
          const result = await mutation<never, User_Create_Input, User>(
            {
              code: "users-create",
            },
            {
              method: "post",
              body: input,
            },
          );

          // update state with response from BE
          setLocalUser(result.data);
        } catch (e: unknown) {
          const err = e as {message: string}
          // if error, clear local state
          setLocalUser(null);
          // and set error
          setError(err.message)
        }

        // update loading status
        setLocalUserStatus("idle");
      };

      // perform request at the same time
      await Promise.any([updateLocalData(), performRequest()]);
    },
    [debug, mutation, setError, setLocalUser, setLocalUserStatus],
  );

  return {
    userCreate,
  };
};

export { useUserCreate };
