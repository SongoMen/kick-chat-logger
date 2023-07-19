
<span class="loader"></span>

<style lang="scss">
    .loader {
        width: 25px;
        height: 25px;
        border: 4px solid $grey;
        border-bottom-color: transparent;
        border-radius: 50%;
        display: inline-block;
        box-sizing: border-box;
        animation: rotation 1s linear infinite;
    }

    @keyframes rotation {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    } 
</style>