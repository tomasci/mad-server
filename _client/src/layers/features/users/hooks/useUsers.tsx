import { useCallback } from "react";
import { User_Create_Input, User_Login_Input } from "../types.ts";
import { useAtom } from "jotai";
import {
  localUserAtom,
  localUserStatusAtom,
  localUserUpdaterAtom,
} from "../store.ts";
import { useNetwork } from "../../../network/hooks/useNetwork.tsx";
import { useDebug } from "../../debug/hooks/useDebug.tsx";

const useUsers = () => {
  // props
  const debug = useDebug({ prefix: "useUsers" });
  const { mutation } = useNetwork();

  // state
  const [localUser] = useAtom(localUserAtom);
  const [, setLocalUser] = useAtom(localUserUpdaterAtom);
  const [localUserStatus, setLocalUserStatus] = useAtom(localUserStatusAtom);

  // functions
  const userCreate = useCallback(async (input: User_Create_Input) => {
    debug.log("userCreate 0");

    const updateLocalData = async () => {
      debug.log("userCreate updateLocalData 0");

      setLocalUser({
        username: input.username,
        email: input.email,
      });
    };

    const performRequest = async () => {
      debug.log("userCreate performRequest 0");
      setLocalUserStatus("loading");

      try {
        debug.log("userCreate performRequest try mutation");
        const result = await mutation();
        if (Math.random() > 0.75) {
          throw new Error("signup failed");
        }
      } catch (e) {
        debug.error("userCreate performRequest error", e);
        setLocalUser(null);
      }

      setLocalUserStatus("idle");
    };

    await Promise.any([updateLocalData(), performRequest()]);
  }, []);

  const userLogin = useCallback(async (input: User_Login_Input) => {
    return;
  }, []);

  return {
    userCreate,
    userLogin,
    localUser,
    localUserStatus,
  };
};

export { useUsers };
