import { useCallback, useMemo } from "react";
import { useAtom } from "jotai";
import {
  feature1SomethingAtom,
  feature1SomethingRequestStatusAtom,
} from "../store.ts";
import { useNetwork } from "../../../network/hooks/useNetwork.tsx";

const useFeature1 = () => {
  // props
  const { addCallback, removeCallback, query, mutation } = useNetwork();

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
  const onResponse = useCallback((featureMethod: string, result: any) => {
    console.log("useFeature1 onResponse", featureMethod, result);

    setSomething((prev) => {
      console.log("something prev", prev);
      return prev;
    });

    switch (featureMethod) {
      case "something":
        // todo: pass result to
        // doSomething_onResponse()
        // todo: then in this function, decide what to do with response according to status
        // setSomething(result.response)
        // setSomethingStatus(result.status)
        // for example, if it was mutation, and it tried to save something on backend, and BE returned result, you maybe need to update
        // already existing record on frontend
        // or maybe it returned error, and you need no mark it as failed etc
        break;

      case "something2":
        break;
    }
  }, []);

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

    // todo: another idea is to add callbacks directly to requests (in query and mutation) and remove addCallback
    // it will look like
    // const perform = async () => {
    //     query(config, (result) => {
    //         // callback here
    //         // try use state here (without setSomething(prev)), just "something"
    //     })
    // }

    await Promise.any([updateLocalData(), query(config.name)]);
  }, [config, setSomething, something, query]);

  const init = useCallback(() => {
    console.log("useFeature1 init", config);
    addCallback(config.name, (featureMethod: string, result: any) => {
      onResponse(featureMethod, result);
    });

    return () => {
      removeCallback(config.name);
    };
  }, [config, addCallback, onResponse]);

  // effects
  // useLayoutEffect(() => {
  //     addCallback(config.name, (featureMethod: string, result: any) => {
  //         onResponse(featureMethod, result)
  //     })
  // }, [config, onResponse])

  return {
    config,
    onResponse,
    init,
    doSomething,
    something,
    somethingStatus,
  };
};

export { useFeature1 };
