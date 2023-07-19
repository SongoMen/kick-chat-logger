<script lang="ts">
    import { onMount } from "svelte";
    import { makeAPIRequest } from "$lib/Requests"
    import Loader from "$components/global/Loader.svelte"
    export let defaultChannel: string | null = "", defaultUser: string | null = "", searchHandler: ((user: string, channel: string, period: string) => void) | null = null;

    let channels: string[] = [];
    let selectedChannel: string = "";
    let username: string = defaultUser || "";
    let loaded: boolean = false;

    const handleSearch = () => {
        if (searchHandler) {
            window.history.pushState({}, '', `/logs?channel=${selectedChannel}&user=${username}`);
            searchHandler(username, selectedChannel, "");
        } else window.location.href = `/logs?channel=${selectedChannel}&user=${username}`
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
        <Loader/>
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