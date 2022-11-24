// const baseURL = 'https://api.sampleapis.com/coffee/hot';
  const baseURL = 'https://c269-2001-fb1-16a-470-5da5-6978-f8b3-a919.ap.ngrok.io/course';
fetch(baseURL)
.then(resp => resp.json())
.then(data => appendData(data))
.catch(err => console.log(`error: `+ err));

function appendData(data) {
  var mainContainer = document.getElementById("myData");
  for (var i = 0; i < data.length; i++) {
      var div = document.createElement("div");
      div.innerHTML = `CourseID: ${data[i].ID} ${data[i].Name} ${data[i].Price} ${data[i].Instructor}`;
      mainContainer.appendChild(div);
  } 
}