import { useCallback } from "react";
import { useAtom } from "jotai";
import { networkCallbackListAtom } from "../store.ts";

const useNetwork = () => {
  // props

  // state
  const [callbackList, setCallbackList] = useAtom(networkCallbackListAtom);

  // functions
  // const whenRequestEnded = useCallback(
  //   (featureName: string, featureMethod: string) => {
  //     const result = {
  //       response: {},
  //       status: {},
  //     };
  //
  //     switch (featureName) {
  //       case "feature1":
  //         break;
  //     }
  //   },
  //   [],
  // );

  const addCallback = useCallback(
    (
      featureName: string,
      callback: (featureName: string, result: any) => void,
    ) => {
      console.log("useNetwork addCallback", featureName, callback);

      setCallbackList((prev) => {
        prev[featureName] = callback;
        return prev;
      });

      return;
    },
    [setCallbackList],
  );

  const removeCallback = useCallback((featureName: string) => {
    console.log("useNetwork removeCallback", featureName);
    setCallbackList((prev) => {
      prev[featureName] = null;
      return prev;
    });
  }, []);

  const query = useCallback(async (featureName: string) => {
    setTimeout(() => {
      console.log("useNetwork query", featureName);
      // when request ended -> call callback by feature name

      const c = callbackList[featureName];
      if (c) {
        c("method placeholder", "any result");
      }

      return;
    }, 1000);
  }, []);

  const mutation = useCallback(async () => {
    return;
  }, []);

  return {
    addCallback,
    removeCallback,
    query,
    mutation,
  };
};

export { useNetwork };
