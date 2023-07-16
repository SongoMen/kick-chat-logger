<script lang="ts">
    import { onMount } from "svelte";
    import { makeAPIRequest } from "$lib/Requests"
    export let defaultChannel: string | null = "", defaultUser: string | null = "", searchHandler: (user: string, channel: string, period: string) => void = () => {};

    let channels: string[] = [];
    let selectedChannel: string = "";
    let username: string = defaultUser || "";
    let loaded: boolean = false;

    const handleSearch = () => {
        window.history.pushState({}, '', `/logs?channel=${selectedChannel}&user=${username}`);
        if (searchHandler) searchHandler(username, selectedChannel, "");
    }

    onMount(async () => {
        channels = await makeAPIRequest("/channels") || [];
        if (defaultChannel && (channels || []).includes(defaultChannel)) selectedChannel = defaultChannel;
        else selectedChannel = (channels || [])[0] || "";
        loaded = true;
    });
</script>

<div class="query">
    {#if !loaded}
        <div>Loading...</div>
    {:else}
        <form on:submit|preventDefault={handleSearch}>
            <select bind:value={selectedChannel}>
                {#each channels as channel}
                    <option>{channel}</option>
                {/each}
            </select>
            <input bind:value={username} type="text" placeholder="Username">
            <a class="button" href="/logs?channel={selectedChannel}&user={username}" on:click={() => searchHandler(username, selectedChannel, "")}>Search</a>
        </form>
    {/if}
</div>