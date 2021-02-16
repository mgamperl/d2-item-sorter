import logo from "./logo.svg";
import "./App.css";

import { useEffect, useState } from "react";

const POSTRequestOptions = {
  method: "POST",
  headers: { "Content-Type": "application/json" },
};

function App() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [sortingSharedStash, setSortingSharedStash] = useState(false);
  const [recorderRunning, setRecorderRunning] = useState(false);

  useEffect(() => {
    // GET request using fetch inside useEffect React hook
    setLoading(true);
    fetch("/api/grailStatus")
      .then((response) => response.json())
      .then((data) => {
        setItems(data);
        setLoading(false);
      });

    // empty dependency array means this effect will only run once (like componentDidMount in classes)
  }, []);

  function startRunRecorder() {
    setRecorderRunning(true)
    fetch("/api/startRunRecorder",POSTRequestOptions)
      .then((response) => response.json())
      .then((data) => {
        setRecorderRunning(true)
      });
  }

  function stopRunRecorder() {
    fetch("/api/stopRunRecorder",POSTRequestOptions)
      .then((response) => response.json())
      .then((data) => {
        setRecorderRunning(false)
      });
  }

  function reloadData() {
    setLoading(true);
    fetch("/api/grailStatus")
      .then((response) => response.json())
      .then((data) => {
        setItems(data);
        setLoading(false);
      });
  }

  function sortSharedStash() {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
    };

    setSortingSharedStash(true);
    fetch("/api/sharedStash/sort", requestOptions)
      .then((response) => response.json())
      .then(() => setSortingSharedStash(false));
  }

  const uniqueMaxCount = items.filter((item) => item.isUnique).length
  const uniqueCount = items.filter((item) => item.isUnique && item.count > 0).length

  const setMaxCount = items.filter((item) => item.isSet).length
  const setCount = items.filter((item) => item.isSet && item.count > 0).length

  const runeMaxCount = items.filter((item) => item.isRune).length
  const runeCount = items.filter((item) => item.isRune && item.count > 0).length

  const overallMaxCount = uniqueMaxCount + setMaxCount + runeMaxCount
  const overallCount = uniqueCount + setCount + runeCount


  return (
    <div className="App">
      <div style={{ padding:10, paddingBottom: 50, display: "flex", justifyContent: "flex-start", alignContent: "center", alignItems: "center"}}>
        <div style={{flexGrow: 1, textAlign: "left"}}>
          <div style={{fontSize: "18pt"}}>Holy Grail {loading ? "(loading ...)" : ""}</div>
          <div style={{fontSize: "12pt"}}><Statistic name="Unique Items" count={uniqueCount} max={uniqueMaxCount}></Statistic></div>
           <div style={{fontSize: "12pt"}}><Statistic name="Set Items" count={setCount} max={setMaxCount}></Statistic></div>
           <div style={{fontSize: "12pt"}}><Statistic name="Runes" count={runeCount} max={runeMaxCount}></Statistic></div>
           <div style={{fontSize: "14pt"}}><Statistic name="Overall" count={overallCount} max={overallMaxCount}></Statistic></div>
        </div>
        <div style={{ display: "flex", height:"30px", alignItems :"center", alignContent: "center" }}>
          <button disabled={loading} onClick={() => reloadData()}>
            Reload
          </button>
          <button
            disabled={sortingSharedStash}
            onClick={() => sortSharedStash()}
          >
            Sort Shared Stash
          </button>

            <button
            disabled={!recorderRunning}
            onClick={() => stopRunRecorder()}
            >
              Stop Run Recorder
            </button>

          <button
            disabled={recorderRunning}
            onClick={() => startRunRecorder()}
            >
              Start Run Recorder
            </button>
        </div>
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          alignItems: "flex-start",
          justifyContent: "space-around",
        }}
      >
        <div
          style={{
            display: "flex",
            flex: "1",
            flexWrap: "wrap",
            justifyContent: "center",
          }}
        >
          {items.map((item) => {
            if (item.isUnique && item.quality === "Normal") {
              return <ItemGridEntry key={item.id} item={item} />;
            } else {
              return null;
            }
          })}
        </div>
        <div
          style={{
            display: "flex",
            flex: "1",
            flexWrap: "wrap",
            justifyContent: "center",
          }}
        >
          {items.map((item) => {
            if (item.isUnique && item.quality === "Exceptional") {
              return <ItemGridEntry key={item.id} item={item} />;
            } else {
              return null;
            }
          })}
        </div>
        <div
          style={{
            display: "flex",
            flex: "1",
            flexWrap: "wrap",
            justifyContent: "center",
          }}
        >
          {items.map((item) => {
            if (item.isUnique && item.quality === "Elite") {
              return <ItemGridEntry key={item.id} item={item} />;
            } else {
              return null;
            }
          })}
        </div>
      </div>
    </div>
  );
}

const Statistic = ({name, count, max}) => {
  return (
    <div>{name}: { count } / {max} <span style={{fontSize: 13}}>({Math.round(count/max * 100)}%)</span></div>
  )
}

const ItemGridEntry = ({ item }) => {
  return (
    <div
      key={item.imageId}
      style={{
        display: "flex",
        flexDirection: "column",
        fontSize: 12,
        width: "200px",
        height: "200px",
      }}
    >
      <div style={{ fontSize: 14 }}>{item.name}</div>
      <div>
        {item.typeName} ({item.quality})
      </div>
      <div
        style={{
          width: "100%",
          height: "100%",
          alignItems: "center",
          display: "flex",
          justifyContent: "center",
          alignSelf: "center",
        }}
      >
        <div
          style={{
            width: "80%",
            height: "80%",
            border: item.count <= 0 ? "2px solid #3D0000" : "2px solid #2F3C2A",
            alignItems: "center",
            display: "flex",
            justifyContent: "center",
            alignSelf: "center",
          }}
        >
          <img
            src={
              process.env.PUBLIC_URL + "/assets/items/" + item.imageId + ".png"
            }
          />
        </div>
      </div>
    </div>
  );
};

export default App;
