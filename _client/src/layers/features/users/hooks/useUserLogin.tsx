import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { useNetwork } from "@/src/layers/network/hooks/useNetwork.tsx";
import { useCallback } from "react";
import {
  User_Login_Input,
  User_Login_Response,
} from "@/src/layers/features/users/types.ts";
import { useAtom } from "jotai/index";
import {
  localUserNetworkErrorAtom,
  localUserStatusAtom,
  localUserTokenAtom,
  localUserUpdaterAtom,
} from "@/src/layers/features/users/store.ts";
import { CommonError } from "@/src/layers/shared/types/CommonTypes.ts";

// with-comments: useUserCreate

const useUserLogin = () => {
  // props
  const debug = useDebug();
  const { mutation } = useNetwork();

  // state
  const [, setLocalUser] = useAtom(localUserUpdaterAtom);
  const [, setLocalUserStatus] = useAtom(localUserStatusAtom);
  const [, setError] = useAtom(localUserNetworkErrorAtom);
  const [, setToken] = useAtom(localUserTokenAtom);

  // functions
  const userLogin = useCallback(
    async (input: User_Login_Input) => {
      debug.log("userLogin call");

      const updatedLocalData = async () => {
        setError(null);

        setLocalUser({
          username: input.username,
        });
      };

      const performRequest = async () => {
        setLocalUserStatus("loading");

        try {
          const result = await mutation<
            never,
            User_Login_Input,
            User_Login_Response
          >(
            {
              code: "users-login",
            },
            {
              method: "post",
              body: input,
            },
          );

          setLocalUser(result.data.user);
          // set token
          setToken(result.data.token);
        } catch (e: unknown) {
          const err = e as CommonError;
          setLocalUser(null);
          setError(err.message);
        }

        setLocalUserStatus("idle");
      };

      await Promise.any([updatedLocalData(), performRequest()]);
    },
    [debug, mutation, setError, setLocalUser, setLocalUserStatus, setToken],
  );

  return { userLogin };
};

export { useUserLogin };
