function t() {
    return function () {
        console.log(this);
    }
}

t()();