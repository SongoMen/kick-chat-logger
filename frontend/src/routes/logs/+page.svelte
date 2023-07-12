<script lang="ts">
    import { onMount } from "svelte";
	import { makeAPIRequest } from '$/lib/Requests';
    import { page } from '$app/stores';
    import type { APIUserLog, APILogsResponse } from "$types/logs"
    
    let channel = $page.url.searchParams.get('channel');
    let user = $page.url.searchParams.get('user');
    let loaded = false;
    let logsByPeriod: { [key: string]: APIUserLog[] } = {};

    onMount(async () => {
        const data: APILogsResponse = await makeAPIRequest('/logs?channel='+channel+'&user='+user);
        console.log(data)
        if (!data.error) {
            for (const period of data['periods']) {
                logsByPeriod[period] = [];
            }
            const split = data['messages'][0].date.split('-')
            const period = `${split[0]}-${split[1]}`;
            logsByPeriod[period] = data['messages'];
        }
        loaded = true;
    });
</script>

<div class="logs">
    <div class="logs__query">
        <input type="text" bind:value={channel}>
        <input type="text" bind:value={user}>
        <a class="button" href="/logs?channel={channel}&user={user}">Search</a>
    </div>
    <div class="logs__periods">
        {#if Object.keys(logsByPeriod).length}
            {#each Object.keys(logsByPeriod) as period}
                <div class="logs__period">
                    {period}
                </div>
            {/each}
        {:else if loaded}
            <div>Not found</div>
        {/if}
    </div>
</div>

<style lang="scss">
    .logs {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
</style>