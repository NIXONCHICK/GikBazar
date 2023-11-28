"use strict"

const toggleBtn = document.querySelector('.list-header');
const mainSelector = document.querySelector('.main-selector');

toggleBtn.onclick = function() {
    mainSelector.classList.toggle('show');
};