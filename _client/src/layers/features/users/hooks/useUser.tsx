import { useAtom } from "jotai/index";
import {
  localUserAtom,
  localUserNetworkErrorAtom,
  localUserStatusAtom,
  localUserTokenAtom,
} from "@/src/layers/features/users/store.ts";

const useUser = () => {
  // props
  // const debug = useDebug();

  // state
  const [localUser] = useAtom(localUserAtom);
  const [localUserStatus] = useAtom(localUserStatusAtom);
  const [localUserNetworkError] = useAtom(localUserNetworkErrorAtom);
  const [localUserToken] = useAtom(localUserTokenAtom);

  return {
    localUser,
    localUserStatus,
    localUserNetworkError,
    localUserToken,
  };
};

export { useUser };
