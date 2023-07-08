<script lang="ts">
    import { onMount } from "svelte";
    import { makeAPIRequest } from "$lib/Requests"

    let channels: string[] = [];
    let selectedChannel: string = "";
    let username: string = "";

    onMount(async () => {
        channels = await makeAPIRequest("/channels");
        selectedChannel = channels[0] || "";
    });
</script>

<div class="home-container">
    <div class="buttons">
        <select bind:value={selectedChannel}>
            {#each channels as channel}
                <option>{channel}</option>
            {/each}
        </select>
        <input bind:value={username} type="text" placeholder="Username">
        <button>Search</button>
    </div>
</div>

<style lang="scss">
    .home-container {
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>