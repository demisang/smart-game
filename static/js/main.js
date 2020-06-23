const app = new Vue({
    el: "#app",
    data: {
        message: "Hello Vue!",
        board: {},
        selectedFigure: 0
    },
    methods: {
        selectFigure: function (event) {
            this.selectedFigure = event.target.dataset.figureId;
        },
        moveFigure: function (event) {
            let direction, rotate, flip;
            if (event.target.id === "moveUp") {
                direction = "1"
            } else if (event.target.id === "moveLeft") {
                direction = "2"
            } else if (event.target.id === "moveRight") {
                direction = "3"
            } else if (event.target.id === "moveDown") {
                direction = "4"
            } else if (event.target.id === "rotate") {
                rotate = "1"
            } else if (event.target.id === "flip") {
                flip = "1"
            }

            axios.put("/move", {
                figureId: this.selectedFigure,
                direction: direction,
                rotate: rotate,
                flip: flip,
            }).then(response => {
                this.board = response.data
            })
        },
        solveBoard: function (event) {
            axios.put("/solve").then(response => {
                this.board = response.data
            })
        }
    },
    mounted() {
        axios.get("/board").then(response => {
            this.board = response.data
        })
    }
});

$(document).on("keydown", function (e) {
    if (!e.keyCode) {
        return
    }
    let button;
    if (e.keyCode === 37) {
        button = $("#moveLeft")
    } else if (e.keyCode === 38) {
        button = $("#moveUp")
    } else if (e.keyCode === 39) {
        button = $("#moveRight")
    } else if (e.keyCode === 40) {
        button = $("#moveDown")
    } else if (e.ctrlKey && e.keyCode === 32) {
        button = $("#flip")
    } else if (!e.ctrlKey && e.keyCode === 32) {
        button = $("#rotate")
    }

    if (button) {
        button.click()
    }
});