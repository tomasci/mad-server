import { useAtom } from "jotai/index";
import {
  localUserAtom, localUserNetworkErrorAtom,
  localUserStatusAtom,
} from "@/src/layers/features/users/store.ts";

const useUser = () => {
  // props
  // const debug = useDebug();

  // state
  const [localUser] = useAtom(localUserAtom);
  const [localUserStatus] = useAtom(localUserStatusAtom);
  const [localUserNetworkError] = useAtom(localUserNetworkErrorAtom)

  return {
    localUser,
    localUserStatus,
    localUserNetworkError
  };
};

export { useUser };
