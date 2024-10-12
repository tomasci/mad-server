import { useCallback, useMemo } from "react";
import { useAtom } from "jotai";
import {
  feature1SomethingAtom,
  feature1SomethingRequestStatusAtom,
} from "../store.ts";
import { useNetwork } from "@/src/layers/network/hooks/useNetwork.tsx";

const useFeature1 = () => {
  // props
  const { query, mutation } = useNetwork();

  // state
  const config = useMemo(() => {
    return {
      name: "feature1",
    };
  }, []);

  const [something, setSomething] = useAtom(feature1SomethingAtom);
  const [somethingStatus, setSomethingStatus] = useAtom(
    feature1SomethingRequestStatusAtom,
  );

  // functions
  const doSomething = useCallback(async () => {
    // work local with data
    // update local store
    // make request
    // after performing request it will update local data & store according to result

    const updateLocalData = async () => {
      setSomething(`I started here! > ${new Date().getTime()}`);
      console.log("useFeature1 doSomething updateLocalData");
      return;
    };

    const perform = async () => {
      try {
        const result = await query();
        setSomething((prev) => {
          console.log("doSomething perform", result, prev);
          return prev; // Keeping previous value just to log the right thing
        });
      } catch (error) {
        console.error("Error in perform: ", error);
      }
    };

    await Promise.any([updateLocalData(), perform()]);
  }, [config, setSomething, query]);

  return {
    config,
    doSomething,
    something,
    somethingStatus,
  };
};

export { useFeature1 };
