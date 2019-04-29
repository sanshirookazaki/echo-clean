<template>
    <div>
        <h2>task lists</h2>
        <ul>
            <li v-for="task in tasks" v-bind:key="task.id" v-on:change="toggleTaskStatus(task)">
                   <input type="checkbox" v-bind:checked="task.done">
                   {{task.name}}
            </li>
        </ul>

        <form v-on:submit.prevent="addTask">
            <input type="text" v-model="newTaskName" placeholder="new task">
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newTaskName: "",
        }
    },

    computed: {
        tasks () {
            return this.$store.state.tasks
        },
    },

    methods: {
        addTask() {
            this.$store.commit("addTask", {
                name: this.newTaskName,
            })
            this.newTaskName = ""
        },

        toggleTaskStatus(task) {
            this.$store.commit("toggleTaskStatus", {
                id: task.id
            })
        },
    }
}
</script>