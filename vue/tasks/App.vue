<template>
    <div>
        <h2>task lists</h2>
        <ul>
            <li v-for="task in tasks" v-bind:key="task.id">
                    <input type="checkbox" v-on:change="toggleTaskStatus(task)" v-bind:checked="task.done">
                    {{task.name}}
                    -
                    <span v-for="id in task.labelIds" v-bind:key="id">
                        {{getLabelText(id)}}
                    </span>
            </li>
        </ul>

        <form v-on:submit.prevent="addTask">
            <input type="text" v-model="newTaskName" placeholder="new task">
        </form>

        <h2>Label List</h2>
        <ul>
            <li v-for="label in labels" v-bind:key="label.id">
                <input type="checkbox" v-bind:value="label.id" 
                v-model="newTaskLabelIds">
                    {{label.text}}
            </li>
        </ul>

        <form v-on:submit.prevent="addLabel">
            <input type="text" v-model="newLabelText" placeholder="newlabel">
        </form>

        <h2>Label Filter</h2>
        <ul>
            <li v-for="label in labels" v-bind:key="label.id">
                <input type="radio" v-bind:checked="label.id === filter"
                    v-on:change="changeFilter(filter.id)">
                {{ label.text }}
            </li>

            <li>
                <input type="radio" v-bind:checked="filter === null"
                    v-on:change="changeFilter(null)">
                nofilter
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newTaskName: "",

            newTaskLabelIds: [],

            newLabelText: ""
        }
    },

    computed: {
        tasks () {
            return this.$store.getters.filteredTask
        },

        labels() {
            return this.$store.state.tasks
        },

        filter() {
            return this.$store.state.filter
        }
    },

    methods: {
        addTask() {
            this.$store.commit("addTask", {
                name: this.newTaskName,
                labelIds: this.newTaskLabelIds 
            })
            this.newTaskName = ""
            this.newTaskLabelIds = []
        },

        toggleTaskStatus(task) {
            this.$store.commit("toggleTaskStatus", {
                id: task.id
            })
        },

        addLabel() {
            this.$store.commit("addLabel", {
                text: this.newLabelText
            })
            this.newLabelText = ""
        },

        getLabelText(id) {
            const label = this.labels.filter(label => label.id === id)[0]
            return label ? label.text: ""
        },

        changeFilter (labelId) {
            this.$store.commit('changeFilter', {
                filter: labelId
            })
        },
    }
}
</script>