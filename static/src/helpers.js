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
