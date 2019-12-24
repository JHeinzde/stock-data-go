<template>
    <div id="stock-form">
        <form @submit.prevent="handleSubmit">
            <label> Stock Symbol </label>
            <input ref="first"
                   type="text"
                   :class="{'has-error': submitting && invalidSymbol }"
                   v-model="symbol"
                   @focus="clearStatus">
            <p v-if="error && submitting" class="error-message">
                Please fill out all the required fields
            </p>
            <p v-if="success"
               class="success-message">
                Stock successfully inserted
            </p>
            <p v-if="error && insertError" class="error-message">
                Stock could not be inserted into the database.
            </p>
            <button> Add Symbol</button>
        </form>
    </div>
</template>

<script>
    export default {
        name: "StockForm",
        data() {
            return {
                error: false,
                submitting: false,
                success: false,
                insertError: false,
                symbol: ''
            }
        },
        computed: {
            invalidSymbol() {
                return this.symbol === ''
            }
        },
        methods: {
            handleSubmit() {
                this.clearStatus();
                this.submitting = true;

                if (this.invalidSymbol) {
                    this.error = true;
                    return
                }

                this.$emit('add:symbol', this.symbol);
                this.$ref.first.focus();
                this.symbol = '';
                this.success = false;
                this.error = false;
            },
            clearStatus() {
                this.success = false;
                this.error = false;
                this.insertError = false;
            }
        }
    }
</script>

<style scoped>
    form {
        margin-bottom: 2rem;
    }

    [class*="-message"] {
        font-weight: 500;
    }

    .error-message {
        color: #d33c40;
    }

    .success-message {
        color: #32a95d;
    }
</style>
