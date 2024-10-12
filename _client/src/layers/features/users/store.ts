import { atom } from "jotai";
import { User } from "./types.ts";
import { merge } from "lodash";
import { NetworkEntityStatus } from "../../network/types.ts";

const localUserAtom = atom<User | null>(null);
const localUserUpdaterAtom = atom(null, (get, set, props: User | null) => {
  if (!props) {
    set(localUserAtom, null);
    return;
  }

  const currentLocalUser = get(localUserAtom);
  const updatedLocalUser = merge({}, currentLocalUser, props);
  set(localUserAtom, updatedLocalUser);
});
const localUserStatusAtom = atom<NetworkEntityStatus>("idle");

export { localUserAtom, localUserUpdaterAtom, localUserStatusAtom };
