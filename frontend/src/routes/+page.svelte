<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	const stocksData = writable({});
	const selectedStock = writable('');
	const options = writable({});
	const chart = writable(null);
	const stockFundamentals = writable({});

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

		chart.set(new ApexCharts(document.querySelector('#chart'), $options));
		$chart.render();
	});
</script>

<div class="flex flex-row justify-between w-screen h-screen">
	<div class="flex flex-col w-2/12 h-full">
		{#each Object.keys($stocksData) as stock}
			<button on:click={() => updateOptions(stock)}>{stock}</button>
		{/each}
	</div>
	<div class="flex flex-col w-8/12 h-full">
		<div id="chart" class="w-full m-9" />
		<div>
			<table>
				<th>EBITDA</th>
				<tr>{$stockFundamentals.EBITDA}</tr>
			</table>
		</div>
	</div>
	<div class="flex flex-col w-3/12">news and stuff</div>
</div>
