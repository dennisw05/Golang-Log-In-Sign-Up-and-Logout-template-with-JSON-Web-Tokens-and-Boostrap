const usernameInput = document.getElementById('username-input')
usernameInput.addEventListener('input', () => {
    if (!usernameInput.checkValidity()) {
        usernameInput.classList.add('is-invalid')
    } else {
        usernameInput.classList.remove('is-invalid')
    }
})

const passwordInput = document.getElementById('password-input')
passwordInput.addEventListener('input', () => {
    CheckPasswords()
    if (!passwordInput.checkValidity()) {
        passwordInput.classList.add('is-invalid')
    } else {
        passwordInput.classList.remove('is-invalid')
    }
})

const passwordConfirmInput = document.getElementById('password-confirm-input')
passwordConfirmInput.addEventListener('input', CheckPasswords)

function CheckPasswords() {
    if (passwordConfirmInput.value != passwordInput.value) {
        passwordConfirmInput.classList.add('is-invalid')
    } else {
        passwordConfirmInput.classList.remove('is-invalid')
    }
}