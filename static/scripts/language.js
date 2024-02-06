document.getElementById('lang-toggler').onclick = () =>
    fetch('/toggle-lang').then(() => {
        location.reload();
    });
