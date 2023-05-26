<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	const stocksData = writable({});
	const selectedStock = writable('');
	const options = writable({});
	const chart = writable(null);
	const stockFundamentals = writable({});
	const news = writable([]);

	$: stocksData.subscribe((data) => {
		options.set({
			chart: {
				type: 'candlestick'
			},
			series: [
				{
					name: $selectedStock,
					data: data[$selectedStock]?.values.map((v) => {
						return {
							x: new Date(v.datetime),
							y: [v.open, v.high, v.low, v.close]
						};
					})
				}
			]
		});
	});

	function updateOptions(stock: string) {
		const chart = new ApexCharts(document.querySelector('#chart'), $options);
		chart.render();
		selectedStock.set(stock);
		options.update((o) => {
			return {
				...o,
				series: [
					{
						name: stock,
						data: $stocksData[stock]?.values.map((v) => {
							return {
								x: new Date(v.datetime),
								y: [v.open, v.high, v.low, v.close]
							};
						})
					}
				]
			};
		});
		$chart.updateOptions($options);
	}

	$: selectedStock.subscribe((stock) => {
		fetch('http://localhost:8080/api/v1/stocks/' + stock)
			.then((res) => res.json())
			.then((data) => {
				stockFundamentals.set(data);
			});
	});

	onMount(async () => {
		await fetch('http://localhost:8080/api/v1/users/stocks/daily/1')
			.then((res) => res.json())
			.then((data) => {
				for (let d in data) {
					data[d].values = data[d].values.reverse();
				}

				selectedStock.set(Object.keys(data)[0]);
				stocksData.set(data);
			});

		await fetch('http://localhost:8080/api/v1/stocks/' + $selectedStock)
			.then((res) => res.json())
			.then((data) => {
				console.log(data);
				stockFundamentals.set(data);
			});

		await fetch('http://localhost:8080/api/v1/rss/')
			.then((res) => res.json())
			.then((data) => {
				news.set(data);
			});

		chart.set(new ApexCharts(document.querySelector('#chart'), $options));
		$chart.render();
	});
</script>

<div class="flex flex-row justify-between w-screen h-screen p-4">
	<div class="flex flex-col w-2/12 h-full">
		<h1 class="text-center mb-4 text-xl">Welcome Victor Guidi</h1>
		{#each Object.keys($stocksData) as stock}
			<button
				class="flex flex-col p-2 h-12 w-full border rounded-md items-center justify-center mb-3"
				on:click={() => updateOptions(stock)}>{stock}</button
			>
		{/each}
	</div>
	<div class="flex flex-col w-8/12 h-full">
		<div id="chart" class="w-full m-9" />
		<div class="flex flex-col w-full p-4">
			<h1 class="text-2xl mb-2">Fundamentals</h1>
			<div class="flex w-fulljustify-between">
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">Market Cap</h1>
					<div>{$stockFundamentals.MarketCapitalization}</div>
				</div>
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">EBITDA</h1>
					<div>{$stockFundamentals.EBITDA}</div>
				</div>
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">PERatio</h1>
					<div>{$stockFundamentals.PERatio}</div>
				</div>
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">Market Cap</h1>
					<div>{$stockFundamentals.MarketCapitalization}</div>
				</div>
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">Market Cap</h1>
					<div>{$stockFundamentals.MarketCapitalization}</div>
				</div>
				<div class="border rounded-md ml-2 w-full">
					<h1 class="mb-2">Market Cap</h1>
					<div>{$stockFundamentals.MarketCapitalization}</div>
				</div>
			</div>
		</div>
	</div>
	<div class="flex flex-col w-3/12">
		<div class="flex flex-col h-1/2">Recommendation</div>
		<div class="flex flex-col h-1/2 overflow-auto pr-3">
			<h1 class="mb-2">News</h1>
			{#each $news as n}
				<div class="flex flex-row">
					<div class="flex flex-col p-2 h-max-26 w-full border mb-3 rounded-md">
						<button class="flex flex-col items-start text-left w-full h-full"
							><a href={n?.link}>
								<div class="text-black font-sans text-sm mb-2">{n.title}</div>
								<div class="text-sm">{n.published}</div>
							</a>
						</button>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
