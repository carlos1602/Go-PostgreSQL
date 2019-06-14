const btn = document.getElementById('say1-buton')
btn.addEventListener('click', () => {
  window.location.replace("http://localhost:8081/all-students");
  //alert(' WARNING: ');
})
const btn2 = document.getElementById('say2-buton')
btn2.addEventListener('click', () => {
  //alert(' WARNING: ');
  window.location.replace("http://localhost:8081/student");
})
const btn3 = document.getElementById('say3-buton')
btn3.addEventListener('click', () => {
  window.location.replace("http://localhost:8081/delete-student");
  //alert(' WARNING: ');
})

const btn4 = document.getElementById('say3-buton')
btn4.addEventListener('click', () => {
  window.location.replace("http://localhost:8081/update-student");
  //alert(' WARNING: ');
})
