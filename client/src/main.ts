fetch('http://localhost:1323/products')
  .then((res) => res.json())
  .then((data) => {
    console.log(data);
    document.getElementById('app')!.innerHTML = JSON.stringify(data, null, 2);
  });
