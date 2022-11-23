// const baseURL = 'https://api.sampleapis.com/coffee/hot';
const baseURL = 'https://3707-2001-fb1-16a-470-5da5-6978-f8b3-a919.ap.ngrok.io/course';
fetch(baseURL)
  .then(resp => resp.json())
  .then(data => displayData(data));

function displayData(data) {
  document.querySelector("pre").innerHTML = JSON.stringify(data, null, 2);
}