module.exports = {
    future: {
      removeDeprecatedGapUtilities: true,
      purgeLayersByDefault: true,
      defaultLineHeights: true,
      standardFontWeights: true
    },
    purge: [],
    theme: {
      extend: {
        backgroundImage: theme => ({
          "version-control": "url('vc.jpeg')",
        })
      }
    },
    variants: {},
    plugins: []
  }
  