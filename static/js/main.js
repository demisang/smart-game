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
            let direction;
            if (event.target.id == "moveUp") {
                direction = "up"
            } else if (event.target.id == "moveLeft") {
                direction = "left"
            } else if (event.target.id == "moveRight") {
                direction = "right"
            } else if (event.target.id == "moveDown") {
                direction = "down"
            }


            axios.put("/move", {
                figureId: this.selectedFigure,
                direction: direction
            }).then(response => {
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
