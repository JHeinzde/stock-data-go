<template>
    <div id="app">
        <header id="app-header">Active Imports</header>
        <StockList :stocks="stocks"/>
        <StockForm @add:symbol="addSymbol"/>
    </div>
</template>

<script>
    import StockList from './components/StockList.vue'
    import axios from 'axios'
    import StockForm from "./components/StockForm";

    export default {
        name: 'app',
        components: {
            StockForm,
            StockList
        }, data() {
            return {
                stocks: []
            }
        }, mounted() {
            this.getStocks()
        },
        methods: {
            getStocks() {
                axios.get("http://localhost:80/stocks")
                    .then(response => (this.stocks = response.data))
            },

            async addSymbol(symbol) {
                axios.post("http://localhost:80/stocks", {symbol: symbol})
                    .catch(error => {
                        if (error.response.status === 500) {
                            this.StockForm.error = true;
                            this.StockForm.insertError = true;
                        }
                    })
            }
        }
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }

    #app-header {
        font-size: 25px;
    }
</style>
