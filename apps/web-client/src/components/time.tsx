import { useEffect, useState } from "react";
import { DateTime } from "luxon"


function TimeDelta({ dt }: {dt: DateTime}) {
    const compute = () => dt.toRelative();
    const [str, setStr] = useState(compute());
  
    useEffect(() => {
      const handler = () => setStr(compute());
      const timer = setInterval(handler, 1000);
      return () => clearInterval(timer);
    }, [dt]);
  
    return <>{str}</>;
  }
  
  export default TimeDelta;