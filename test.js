const solution = (record) => {
  let arr = record.map((v, i) => v.split(" "));
  let obj = Object.fromEntries(arr);
  console.log(obj);

  for (let i = 0; i < arr.length; i++) {
    let key = arr[i][0];
    console.log(key);
    switch (key) {
      case "Enter":
        return arr[i][2] + " came in";
      // case "Leave":
      //     return arr[i][1] === arr[i][1] ? ""

      default:
        break;
    }
  }
};

solution([
  "Enter uid1234 Muzi",
  "Enter uid4567 Prodo",
  "Leave uid1234",
  "Enter uid1234 Prodo",
  "Change uid4567 Ryan",
]);
