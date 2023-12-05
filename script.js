const $ = q => document.querySelector(q);

const dropdown = $('#nav__dropdown');
const list = $('#nav__links');

dropdown.onclick = () => {
    list.classList.toggle('hide');
};

document.querySelectorAll('.nav__link').forEach(
    e =>
        (e.onclick = () => {
            list.classList.toggle('hide');
            console.log('yes');
        })
);

const goToTop = $('#go-to-up');

goToTop.onclick = () => {
    window.scroll({ top: 0, left: 0, behavior: 'smooth' });
};

window.onscroll = () => {
    if (window.scrollY <= 100) {
        //user is at the top of the page; no need to show the back to top button
        goToTop.style.display = '';
    } else {
        goToTop.style.display = 'block';
    }
};
