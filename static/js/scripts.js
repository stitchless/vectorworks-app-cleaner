// (function(window, document) {
    'use strict';
    document.addEventListener('DOMContentLoaded', () => {
        // const buttons = document.querySelectorAll('i')
        // for (const button of buttons) {
        //     button.addEventListener('click', (event) => {
        //         const target = event.target
        //         console.log(target)
        //         const form = target.closest('form')
        //         console.log(form.target)
        //     });
        // }
        const buttons = document.getElementsByClassName('fas');
        for (const button of buttons) {
            const form = button.nextElementSibling
            button.addEventListener('click', (event) => {
                form.submit()
            });
        }
    })
// })(window, document);