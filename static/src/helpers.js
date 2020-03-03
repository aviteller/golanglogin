import Cookies from "./Cookies";
let c = new Cookies();

export let cookieUser = c.getCookie("jwt")
  ? {
      id: JSON.parse(c.getCookie("jwt")).id,
      name: JSON.parse(c.getCookie("jwt")).name,
      token: JSON.parse(c.getCookie("jwt")).token
    }
  : {};

export const mealTypeToString = mealType => {
  let mealReturn = mealTypeOptions.filter(meal => meal.id == mealType);
  return mealReturn[0].text;
};

export const mealTypeOptions = [
  { id: 0, text: "Breakfast" },
  { id: 1, text: "Lunch" },
  { id: 2, text: "Supper" },
  { id: 3, text: "Snack" }
];

export const genderOptions = [
  { id: "Male", text: "Male" },
  { id: "Female", text: "Female" }
];

export const convertDateToAge = (dateString, returnString = true) => {
  if (!dateString) return 0;
  let firstDate = new Date(dateString);
  let diff = new Date(new Date() - firstDate);
  if (returnString) {
    return `${diff.toISOString().slice(0, 4) -
      1970} years and ${diff.getMonth()} months`;
  } else {
    return diff;
  }
};

export const convertDateToString = date => {
  let options = {
    weekday: "short",
    year: "numeric",
    month: "short",
    day: "numeric"
  };

  let dateThing = new Date(date);

  return dateThing.toLocaleDateString("en-UK", options);
};

export const waitFor = ms => new Promise(s => setTimeout(s, ms));