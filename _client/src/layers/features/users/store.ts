import { atom } from "jotai";
import { User, User_Token } from "./types.ts";
import { NetworkEntityStatus } from "@/src/layers/network/types.ts";
import { merge } from "lodash";
import { atomWithStorage } from "jotai/utils";

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
const localUserNetworkErrorAtom = atom<string | null>(null);
const localUserTokenAtom = atomWithStorage<User_Token | null>("token", null);

export {
  localUserAtom,
  localUserUpdaterAtom,
  localUserStatusAtom,
  localUserNetworkErrorAtom,
  localUserTokenAtom,
};
