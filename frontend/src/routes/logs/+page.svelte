<script lang="ts">
    import { onMount } from "svelte";
    import { page } from '$app/stores';
    import Loader from "$components/global/Loader.svelte"

    import type { APIUserLog, APILogsResponse, APIChannelInfo } from "$types/logs"

	import { makeAPIRequest } from '$/lib/Requests';
    import { getAllEmotes } from '$/lib/emotes';
    import MessagesContainer from "$components/logs/MessagesContainer.svelte"
    import LogsQuery from "$components/logs/LogsQuery.svelte"

    let loadedData: Record<string, string> = {'channel': $page.url.searchParams.get('channel') || '', 'user': $page.url.searchParams.get('user') || ''};
    let loaded = false;
    let channelInfo: APIChannelInfo;
    let emotes: Record<string, string> = {};
    let periods: string[] = [];
    let logsByPeriod: { [key: string]: APIUserLog[] } = {};

    const loadChannelData = async (channel: string) => {
        channelInfo = await makeAPIRequest(`/channel-info?channel=${channel}`);
        emotes = await getAllEmotes(channelInfo['stvID'] || "");
    }

    const retrieveMessages = async (user: string, channel: string, period: string) => {
        if (loaded && (loadedData['channel'] !== channel || loadedData['user'] !== user)) {
            logsByPeriod = {};
        } else if (logsByPeriod[period] && logsByPeriod[period].length && period !== periods[0]) {
            return;
        } else if (!loaded || channel !== loadedData['channel']) {
            await loadChannelData(channel);
        }
        loaded = false;
        let url = `/logs?channel=${channel}&user=${user}`;
        if (period !== "") url += `&period=${period}`;
        const data: APILogsResponse = await makeAPIRequest(url);
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
    }

    onMount(async () => {
        if (loadedData['channel'] && loadedData['user']) retrieveMessages(loadedData['user'], loadedData['channel'], "");
        else loaded = true;
    });
</script>

<div class="logs">
    <LogsQuery defaultChannel={loadedData['channel']} defaultUser={loadedData['user']} searchHandler={retrieveMessages} />
    <div class="logs__periods">
        {#if Object.keys(logsByPeriod).length}
            {#each periods as period}
                <MessagesContainer channel={loadedData['channel']} getMessages={retrieveMessages} emotes={emotes} username={loadedData['user']} label={period} messages={logsByPeriod[period]} on:click={() => retrieveMessages(period)}/>
            {/each}
        {:else if loaded && loadedData['user']}
            <h2>No messages found.</h2>
        {:else if !loaded}
            <Loader/>
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