var x = document.querySelector('#t')
x.onclick = open;
function open(){
    document.getElementById('header').classList.toggle('active');
    document.getElementById('nav').classList.toggle('active');
    document.getElementById('t').classList.toggle('is-active');
    document.body.classList.toggle('open');
}

function hideMsg() {
    document.getElementById('login-msg').classList.add('hidden');
}
