import { atom } from "jotai";

const networkCallbackListAtom = atom<{
  [featureName: string]: ((featureMethod: string, result: any) => void) | null;
}>({});

export { networkCallbackListAtom };
