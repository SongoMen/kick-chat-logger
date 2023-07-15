<script lang="ts">
    import { onMount } from "svelte";
	import { makeAPIRequest } from '$/lib/Requests';
    import { getGlobalEmotes } from '$/lib/7tv';
    import { page } from '$app/stores';
    import type { APIUserLog, APILogsResponse } from "$types/logs"
    import MessagesContainer from "$components/logs/MessagesContainer.svelte"
    import LogsQuery from "$components/logs/LogsQuery.svelte"

    let loadedData: Record<string, string> = {'channel': $page.url.searchParams.get('channel') || '', 'user': $page.url.searchParams.get('user') || ''};
    let loaded = false;
    let emotes: Record<string, string> = {};
    let periods: string[] = [];
    let logsByPeriod: { [key: string]: APIUserLog[] } = {};

    const retrieveMessages = async (user: string, channel: string, period: string) => {
        console.log(loadedData, loaded, 'u', user, 'c', channel, 'p', period);

        if (loaded && (loadedData['channel'] !== channel || loadedData['user'] !== user)) {
            logsByPeriod = {};
        } else if (logsByPeriod[period] && logsByPeriod[period].length && period !== periods[0]) {
            return;
        }
        loaded = false;
        let url = `/logs?channel=${channel}&user=${user}`;
        if (period !== "") url += `&period=${period}`;
        const data: APILogsResponse = await makeAPIRequest(url);
        console.log(data)
        if (!data.error && data['messages'].length) {
            for (const period of data['periods']) {
                if (!logsByPeriod[period]) logsByPeriod[period] = [];
            }
            const split = data['messages'][0].date.split('-')
            const period = `${split[0]}-${split[1]}`;
            logsByPeriod[period] = data['messages'];
            if (!periods.length) periods = data['periods'];
        }
        loadedData = { 'channel': channel, 'user': user };
        loaded = true;
        console.log(logsByPeriod)
    }

    onMount(async () => {
        emotes = await getGlobalEmotes()
        await retrieveMessages(loadedData['user'], loadedData['channel'], "");
    });
</script>

<div class="logs">
    <LogsQuery defaultChannel={loadedData['channel']} defaultUser={loadedData['user']} searchHandler={retrieveMessages} />
    <div class="logs__periods">
        {#if Object.keys(logsByPeriod).length}
            {#each periods as period}
                <MessagesContainer channel={loadedData['channel']} getMessages={retrieveMessages} emotes={emotes} username={loadedData['user']} label={period} messages={logsByPeriod[period]} on:click={() => retrieveMessages(period)}/>
            {/each}
        {:else if loaded}
            <h2>No messages found.</h2>
        {/if}
    </div>
</div>

<style lang="scss">
    .logs {
        display: flex;
        flex-direction: column;
        align-items: center;
        &__periods {
            width: 100%;
            display: flex;
            align-items: center;
            flex-direction: column;
            margin-top: 20px;
        }
    }
</style>