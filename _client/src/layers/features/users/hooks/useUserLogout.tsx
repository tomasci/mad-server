import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { useCallback } from "react";
import { useAtom } from "jotai/index";
import {
  localUserNetworkErrorAtom,
  localUserStatusAtom,
  localUserTokenAtom,
  localUserUpdaterAtom,
} from "@/src/layers/features/users/store.ts";

const useUserLogout = () => {
  // props
  const debug = useDebug();

  // state
  const [, setLocalUser] = useAtom(localUserUpdaterAtom);
  const [, setLocalUserStatus] = useAtom(localUserStatusAtom);
  const [, setError] = useAtom(localUserNetworkErrorAtom);
  const [, setToken] = useAtom(localUserTokenAtom);

  // functions
  const userLogout = useCallback(() => {
    debug.log("userLogout call");

    // just reset everything user-related
    setLocalUser(null);
    setLocalUserStatus("idle");
    setError(null);
    setToken(null);
  }, [debug, setError, setLocalUser, setLocalUserStatus, setToken]);

  return { userLogout };
};

export { useUserLogout };
