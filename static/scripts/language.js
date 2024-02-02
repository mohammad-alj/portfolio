let btn = document.getElementById('lang-toggler');
btn.onclick = () => {
    fetch('/toggle-lang')
        .then(() => {
            location.reload();
        })
        .catch(err => console.log(err));
};
