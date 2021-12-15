// Cookie reading function from https://www.w3schools.com/js/js_cookies.asp
function setCookieWithMonthExpiry(name, value) {
    let expiry = new Date();
    expiry.setMonth(expiry.getMonth() + 1); // Adds one month to current date
    document.cookie = `${name}=${value};expires=${expiry.toUTCString()};`;
}

function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = document.cookie;
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function checkCookie(cookieName) {
    let cookieVal = getCookie(cookieName);
    return cookieVal !== "";
}

export {
    setCookieWithMonthExpiry,
    getCookie,
    checkCookie,
}